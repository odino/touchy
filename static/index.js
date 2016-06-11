import React from 'react';
import ReactDOM from 'react-dom';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import RaisedButton from 'material-ui/RaisedButton';
import AppBar from 'material-ui/AppBar';
import GridList from 'material-ui/GridList';

const App = () => (
  <MuiThemeProvider muiTheme={getMuiTheme()}>
    <Root />
  </MuiThemeProvider>
);

function press(key) {
  return () => {
    fetch('/press/' + key, {
      credentials: 'same-origin'
    })
  }
}

class Button extends React.Component {
  render() {
    return <RaisedButton {...this.props} style={{height: "100px", width: '100%'}} labelStyle={{fontSize: "80px"}} onClick={press(this.props.keypress)}/>
  }
}

class Root extends React.Component {
  render() {
    return (
      <div>
        <AppBar title={"Touchy"} showMenuIconButton={false} />
        <GridList style={{alignItems: 'center', flex: 1, justifyContent: "center", margin: '0em', marginTop: "2em"}} cellHeight={200}>
            <Button label={"↑"} keypress={"up"} />
        </GridList>
        <GridList style={{alignItems: 'center', flex: 1, justifyContent: "center", margin: '0em'}} cellHeight={200}>
            <Button label={"←"} keypress={"left"} />
            <Button label={"→"} keypress={"right"} />
        </GridList>
        <GridList style={{alignItems: 'center', flex: 1, justifyContent: "center", margin: '0em'}} cellHeight={200}>
            <Button label={"↓"} keypress={"down"} />
        </GridList>
      </div>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);
