var React = require('react')
var ReactDOM = require('react-dom')

var Menu = React.createClass({
  render: function () {
    return (
      <div id='menu'>
        <div className='ui secondary menu'>
          <a className='item'>Home</a>
        </div>
      </div>
    )
  }
})

var Login = React.createClass({
  getInitialState: function () {
    return {
      link: '#/signup'
    }
  },

  render: function () {
    return (
      <div id='login-form'>
        <h1> Blu Lite</h1>
        <form id='login' className='ui form' method='post'>
          <div className='field'>
            <input type='text' placeholder='User name'/>
          </div>
          <div className='field'>
            <input type='password' placeholder='Password'/>
          </div>
          <div className='field'>
            <a href={this.state.link}>Sign up</a>
          </div>
          <div className='field'>
            <div className='ui fluid blue submit button' id='submit'>Log In</div>
          </div>
          <div className='forgotLink'>
            <a href='#/trouble'>Having trouble logging in?</a>
          </div>
        </form>
        <div className='g-signin2' data-onsuccess='onSignIn' data-theme='dark'></div>
      </div>
    )
  }
})

var Footer = React.createClass({
  getInitialState: function () {
    var year = new Date().getFullYear()
    return {
      year: year
    }
  },
  render: function () {
    return (
      <div id='footer'>
        <p>&copy;Blu Sigma {this.state.year}</p>
      </div>
    )
  }
})

var BluLite = React.createClass({
  render: function () {
    return (
      <div>
        <Login />
        <Footer />
      </div>
    )
  }
})

ReactDOM.render(
  <BluLite />, document.getElementById('blulite')
)
