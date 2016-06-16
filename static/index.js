import React from 'react';
import ReactDOM from 'react-dom';
import {TransitionMotion, spring} from 'react-motion';

let midPoint = window.innerWidth / 2;

const leavingSpringConfig = {stiffness: 60, damping: 15};
const DIRECTIONS = {
  left: "← ← ← ",
  right: "→ → → ",
}

const App = React.createClass({
  getInitialState() {
    return {
      mouse: [],
      now: 't' + 0,
      from: null,
      to: null,
      text: "← swipe left or right →"
    };
  },

  handleMouseMove({pageX, pageY, clientX}) {
    this.setState(() => {
      return {
        mouse: [pageX - 25, pageY - 25],
        now: 't' + Date.now(),
        to: clientX
      };
    });
  },

  /**
   * Flashes the DOM, meaning that it will
   * quickly fade it to full-white and then
   * back to its original style.
   */
  flash() {
    var elem = ReactDOM.findDOMNode(this);

    window.requestAnimationFrame(function() {
      elem.style.transition = "opacity 300ms";
      elem.style.opacity = 0;

      setTimeout(() => {
        elem.style.transition = "opacity 300ms";
        elem.style.opacity = 1
      }, 300);
    });
  },

  /**
   * Presses the given key via the API.
   *
   * Super error handling. Would you care?
   */
  press(key) {
    fetch(`/press/${key}`, {
      credentials: 'same-origin'
    })
  },

  /**
   * Get the direction of the swipe gesture.
   * If the swipe didn't cross half the screen,
   * we discard it.
   */
  getDirection() {
    if (this.state.from < midPoint && this.state.to > midPoint) {
      return "left"
    } else if (this.state.from > midPoint && this.state.to < midPoint) {
      return "right"
    }
  },

  /**
   * Action that encapsulate all the logic we need
   * to do when a valid swipe happens.
   */
  swipe(direction) {
    this.press(direction)
    this.flash()
    this.setState({...this.getInitialState(), text: DIRECTIONS[direction]})
    window.navigator.vibrate && window.navigator.vibrate(150);
  },

  /**
   * When touch is over, let's figure out if we
   * have to trigger a key press.
   * First we check if this was a single click:
   * if so, screw this. Then we check if it was
   * a swipe and only then we trigger the key
   * representing the direction of the swipe.
   */
  onTouchEnd(e) {
    let singleClick = !this.state.to;

    if (!singleClick) {
      let direction = this.getDirection()

      if (direction) {
        this.swipe(direction)
      }

      e.preventDefault();
    }
  },

  willLeave(styleCell) {
    return {
      ...styleCell.style,
      opacity: spring(0, leavingSpringConfig),
      scale: spring(2, leavingSpringConfig),
    };
  },

  getStyles() {
    let {mouse: [mouseX, mouseY], now} = this.state;

    return (mouseX == null) ? [] : [{
      key: now,
      style: {
        opacity: spring(1),
        scale: spring(0),
        x: spring(mouseX),
        y: spring(mouseY),
      }
    }];
  },

  renderSwiper() {
    return (
      <TransitionMotion willLeave={this.willLeave} styles={this.getStyles()}>
        {circles =>
          <div
            onTouchStart={(e) => this.setState({from: e.touches[0].clientX})}
            onTouchEnd={this.onTouchEnd}
            onTouchMove={(e) => this.handleMouseMove(e.touches[0])}
            className="container">
            {circles.map(({key, style: {opacity, scale, x, y}}) =>
              <div
                key={key}
                className="ball"
                style={{
                  opacity: opacity,
                  scale: scale,
                  transform: `translate3d(${x}px, ${y}px, 0) scale(${scale})`,
                  WebkitTransform: `translate3d(${x}px, ${y}px, 0) scale(${scale})`,
                }} />
            )}
          </div>
        }
      </TransitionMotion>
    )
  },

  render() {
    return (
      <div>
        {this.renderSwiper()}
        <div style={{
          position: "absolute",
          width: "100%",
          height: "10%",
          bottom: 0,
          justifyContent: "center",
          alignItems: "center",
          display: "flex"
        }}>
          <span style={{fontSize: "2em"}}>
            {this.state.text}
          </span>
        </div>
      </div>
    );
  },
});

ReactDOM.render(
  <App />,
  document.getElementById('app')
);
