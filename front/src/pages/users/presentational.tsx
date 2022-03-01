import React from "react";
import {Link} from "react-router-dom";

const Presentational = ():JSX.Element => {
  return (
    <>
      <h2>Users</h2>
      <Link to="/users/1" >/users/1</Link><br/>
      <Link to="/users/2" >/users/2</Link><br/>
      <Link to="/users/3" >/users/3</Link>
    </>
  )
}

export const Component = React.memo(Presentational)
