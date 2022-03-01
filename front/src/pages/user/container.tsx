import React from "react";
import {Component} from "./presentational";
import {useParams} from "react-router-dom";


export const Container = (): JSX.Element => {
  const params = useParams<{
    userID: string | undefined
  }>()
  return (
    <Component userID={params.userID ?? ""} />
  )
}
