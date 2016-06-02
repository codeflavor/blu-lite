var React = require('react')
var ReactDOM = require('react-dom')

var App = React.createClass({
  render: function () {
    return (
      <div className="ui input">
        <input type="text" placeholder="Username..."></input>
      </div>
    )
  }
})

ReactDOM.render(
  <App/>, document.getElementById('content')
)
