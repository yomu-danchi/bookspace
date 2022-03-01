import React from "react";

const Presentational = ():JSX.Element => {
  return (
      <h2>Users</h2>
  )
}

export const Component = React.memo(Presentational)
