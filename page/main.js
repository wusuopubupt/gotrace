var scene, camera, renderer, controls;

init();
animate();

function init() {
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
	//camera = new THREE.OrthographicCamera( width / - 2, width / 2, height / 2, height / - 2, -1000, 2000 );
	camera = new THREE.PerspectiveCamera(75, width / height, 1, 1000 );
	camera.position.z = 150;
	camera.updateProjectionMatrix();
	
	mat1 = new THREE.LineBasicMaterial( { color: 0x0000ff, linewidth: 4, } );
	trace = new GoThree.Trace();
	trace.init(scene, data, params);
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
	renderer = new THREE.WebGLRenderer({ alpha: true });
	renderer.setSize( width, height );
	renderer.setClearColor( '#2F2933', 1);

	// leap camera controls
	controls = new THREE.LeapMyControls( camera , controller, renderer.domElement );
	//controls = new THREE.LeapPointerControls( camera , controller, renderer.domElement );

	// CONTROLS
	var orbit = new THREE.OrbitControls( camera, renderer.domElement );
	orbit.addEventListener( 'change', function() {
		trace.onControlsChanged(orbit.object);
	});

	// ADD CUSTOM KEY HANDLERS
	document.addEventListener("keydown", function(event) {trace.Keydown(event)}, false);

	document.body.appendChild( renderer.domElement );

	controller.connect();
}

function animate() {
	controls.update();
	trace.animate();

	requestAnimationFrame(animate);
	stats.begin();
	renderer.render(scene, camera);
    stats.end();
}
