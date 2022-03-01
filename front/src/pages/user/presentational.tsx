import React from "react"

type Props = {
    userID: string
}

const Presentational = (props: Props):JSX.Element => {
  return (
      <h2>User {props.userID }</h2>
  )
}

export const Component = React.memo(Presentational)
