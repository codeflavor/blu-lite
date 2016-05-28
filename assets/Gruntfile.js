module.exports = function(grunt) {
  require('load-grunt-tasks')(grunt);
  var timer = require("grunt-timer");

  my_src_files = ['src/*.js']
  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      build: {
        src: 'src/<%= pkg.name %>.js',
        dest: 'build/<%= pkg.name %>.min.js'
      }
    }
  });

  // Build task
  grunt.registerTask('build', [
    'eslint',
    'jshint',
    'uglify'
  ]);

  grunt.registerTask('default','');
};
