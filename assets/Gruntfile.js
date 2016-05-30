'use strict'

module.exports = function (grunt) {
  require('time-grunt')(grunt)
  require('load-grunt-tasks')(grunt)

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    eslint: {
      target: ['src/*']
    },
    jshint: {
      options: {
        reporter: require('jshint-stylish')
      },
      target: ['file.js']
    },
    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      build: {
        src: 'src/<%= pkg.name %>.js',
        dest: 'build/<%= pkg.name %>.min.js'
      }
    },
    clean: {
      build: {
        src: [
          'build/*'
        ]
      }
    }
  })

  grunt.registerTask('default', [
    'eslint',
    'jshint'
  ])

  grunt.registerTask('install', [
    'bower_install,'
  ])

  // Build task
  grunt.registerTask('build', [
    'eslint',
    'jshint',
    'babel',
    'uglify'
  ])
}
