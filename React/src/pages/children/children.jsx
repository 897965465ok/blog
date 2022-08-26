import React, { useEffect } from "react";
function FancyBorder(props) {
  return (
    <div className={'FancyBorder FancyBorder-' + props.color}>
      { props.children}
      { console.log(props)}
    </div>
  );
}



function WelcomeDialog() {
  return (
    <FancyBorder color="blue">
      <h1 className="Dialog-title">
        Welcome
      </h1>
    </FancyBorder>
  );
}

export default  WelcomeDialog 