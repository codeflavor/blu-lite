var React = require('react')
var ReactDOM = require('react-dom')

var BluLIte = React.createClass({
  render: function () {
    return (
      <h1>Hello World!</h1>
    )
  }
})

ReactDOM.render(
  <BluLIte />, document.getElementById('test')
)
