var source = "static_source";
var pub = "public";
var tmp = "tmp";

module.exports = {
    build_lib_js: {
        filename: 'lib.min.js',
        paths: {
            bowerDirectory: source + '/bower_components',
            bowerrc: source + '/.bowerrc',
            bowerJson: source + '/bower.json'
        },
        dest: source + '/js'
    },
    build_coffee_js: {
        filename: 'app.min.js',
        source: [
            source + "/app/app.coffee",
            source + "/app/services/**/*.coffee",
            source + "/app/filters/**/*.coffee",
            source + "/app/directives/**/*.coffee",
            source + "/app/controllers/**/*.coffee"
        ],
        watch: source + "/app/**/*.coffee",
        dest: source + '/js'
    },
    build_lib_css: {
        filename: 'lib.min.css',
        paths: {
            bowerDirectory: source + '/bower_components',
            bowerrc: source + '/.bowerrc',
            bowerJson: source + '/bower.json'
        },
        dest: source + '/css'
    },
    build_less: {
        filename: 'app.min.css',
        source: [
            source + '/less/app.less'
        ],
        dest: source + '/css',
        watch: source + '/less/**/*.less'
    },
    copy: {
        source: [
            source + '/bower_components/font-awesome/fonts/**/*'
        ],
        dest: source,
        source_dir: source + '/bower_components/font-awesome'
    }
};
