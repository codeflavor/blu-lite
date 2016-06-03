var React = require('react')
var ReactDOM = require('react-dom')

var BluLite = React.createClass({
  render: function () {
    return (
      <div>
        <Menu/>
        <Login/>
      </div>
    )
  }
})

var Menu = React.createClass({
  render: function () {
    return (
      <div className='ui secondary menu'>
        <a className='item'>First Item</a>
      </div>
    )
  }
})

var Login = React.createClass({
  render: function () {
    return (
      <form id='login-form' className='ui form'>
        <div className='field'>
          <input type='text' placeholder='User name'/>
        </div>
        <div className='field'>
          <input type='text' placeholder='Password'/>
        </div>
        <div className='ui fluid blue submit button'>Submit</div>
      </form>
    )
  }
})

ReactDOM.render(
  <BluLite/>, document.getElementById('blulite')
)
