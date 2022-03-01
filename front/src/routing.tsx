import {Route, Routes} from "react-router-dom";
import Users from "./pages/users";
import User from "./pages/user";

const Routing: React.VFC<unknown> = () => (
  <>
    <Routes>
      <Route path="/" element={<Users />} />
      <Route path="/users/:userID"
             element={
               <User />
             }
      />
    </Routes>
  </>
)

export default Routing
