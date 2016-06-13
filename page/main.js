var scene, camera, renderer, controls, orbit, trace, cameraControls;

// leap motion helpers
var baseBoneRotation = ( new THREE.Quaternion ).setFromEuler( new THREE.Euler( 0, 0, Math.PI / 2 ) );
var armMeshes = [];
var boneMeshes = [];

init();
Leap.loop(function(frame){
  // maybe some modifications to your scene
  // ...

  cameraControls.update(frame); // rotating, zooming & panning
  trace.animate();
  renderer.render(scene, camera);
});

animate();

function init() {
	var scale = 5;

	// STATS
	stats = new Stats();
	stats.setMode( 0 ); // 0: fps, 1: ms, 2: mb
	stats.domElement.style.position = 'absolute';
	stats.domElement.style.left = '0px';
	stats.domElement.style.top = '0px';
	document.body.appendChild( stats.domElement );

	controller = new Leap.Controller();
	scene = new THREE.Scene();

	// CAMERA
	width = window.innerWidth;
	height = window.innerHeight;
	var center = new THREE.Vector3(60, -50, -10);
	//camera = new THREE.OrthographicCamera( width / - 2, width / 2, height / 2, height / - 2, -1000*scale, 2000*scale );
	camera = new THREE.PerspectiveCamera(75, width / height, 1, 1000 * scale );
	camera.position.z = 100 * scale;
	camera.position.y = 50 * scale;
	camera.updateProjectionMatrix();
	
	mat1 = new THREE.LineBasicMaterial( { color: 0x0000ff, linewidth: 4, } );
	trace = new GoThree.Trace();
	trace.init(scene, data, params, scale);
	/*
	trace.init(scene, data, {
		allCaps: true,
		zoom: 0.6,
		speed: 1.5,
		angle: 45,
		angle2: 90,
		autoAngle: false,
		totalTime: 100,
		distance: 100,
		autoGrow: false,
		distance2: 20
	});
	*/

	// RENDERER
	renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true, });
	renderer.setSize( width, height );
	renderer.setClearColor( '#1D1F17', 1);

	// light for hand
	var light = new THREE.AmbientLight( 0x505050 ); // soft white light
	scene.add( light );

	// leap camera controls
	//controls = new THREE.LeapMyControls( camera , controller, renderer.domElement );
	//controls = new THREE.LeapPointerControls( camera , controller, renderer.domElement );
	cameraControls = new THREE.LeapCameraControls(camera);
	cameraControls.panEnabled      = true;
	cameraControls.panSpeed        = 1.0;
	cameraControls.panHands        = 2;
	cameraControls.panFingers      = [6,12];
	cameraControls.panRightHanded  = true; // right-handed
	cameraControls.panHandPosition = true; // palm position used
	cameraControls.panStabilized   = true; // stabilized palm position used
	
	cameraControls.rotateEnabled         = true;
	cameraControls.rotateHands           = 1;
	cameraControls.rotateSpeed           = 0.8;
	cameraControls.rotateFingers         = [4, 5];
	cameraControls.rotateRightHanded     = true;
	cameraControls.rotateHandPosition    = false;
	cameraControls.rotateStabilized      = false;

	cameraControls.zoomEnabled         = true;
	cameraControls.zoomHands           = [1,2];
	cameraControls.zoomSpeed           = 1;
	cameraControls.zoomFingers         = [2, 3];
	cameraControls.zoomRightHanded     = true;
	cameraControls.zoomHandPosition    = true;
	cameraControls.zoomStabilized      = true;

	// CONTROLS
	orbit = new THREE.OrbitControls( camera, renderer.domElement );
	orbit.autoRotate = false;
	orbit.autoRotateSpeed = 1.0;
	orbit.addEventListener( 'change', function() {
		trace.onControlsChanged(orbit.object);
	});

	// ADD CUSTOM KEY HANDLERS
	document.addEventListener("keydown", function(event) {keydown(event)}, false);

	document.body.appendChild( renderer.domElement );

	controller.connect();
}

function animate() {

	if (orbit.autoRotate) {
		orbit.update();
	};
	//controls.update();
	trace.animate();

	requestAnimationFrame(animate);
	stats.begin();
	renderer.render(scene, camera);
    stats.end();
}

function keydown(event) {
	switch (event.which) {
		case 80: // 'P' - (Un)Pause autoRotate
			toggleAutoRotate();
			break;
		case 82: // 'R' - Reset
			trace.resetTime();
			break;
		case 83: // 'S' - Slower
			trace.slowdown();
			break;
		case 70: // 'F' - Faster
			trace.speedup();
			break;
	}
}

function toggleAutoRotate() {
	orbit.autoRotate = !orbit.autoRotate;
}

function leapAnimate( frame ) {
	var countBones = 0;
	var countArms = 0;
	armMeshes.forEach( function( item ) { scene.remove( item ) } );
	boneMeshes.forEach( function( item ) { scene.remove( item ) } );
	for ( var hand of frame.hands ) {
		for ( var finger of hand.fingers ) {
			for ( var bone of finger.bones ) {
				if ( countBones++ === 0 ) { continue; }
				var boneMesh = boneMeshes [ countBones ] || addMesh( boneMeshes );
				updateMesh( bone, boneMesh );
			}
		}
		var arm = hand.arm;
		var armMesh = armMeshes [ countArms++ ] || addMesh( armMeshes );
		updateMesh( arm, armMesh );
		armMesh.scale.set( arm.width / 4, arm.width / 2, arm.length );
	}

	animate();
}

function addMesh( meshes ) {
	var geometry = new THREE.BoxGeometry( 1, 1, 1 );
	var material = new THREE.MeshNormalMaterial();
	var mesh = new THREE.Mesh( geometry, material );
	meshes.push( mesh );
	return mesh;
}

function updateMesh( bone, mesh ) {
	mesh.position.fromArray( bone.center() );
	mesh.setRotationFromMatrix( ( new THREE.Matrix4 ).fromArray( bone.matrix() ) );
	mesh.quaternion.multiply( baseBoneRotation );
	mesh.scale.set( bone.width, bone.width, bone.length );
	scene.add( mesh );
}
