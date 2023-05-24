import {User} from './types/UserData'
import { ILogIn } from "./types/authData"
import http from "./common"

const logIn = (data: ILogIn) => {
  return http.post<ILogIn>("/login", data)
}

const AuthService = {
  logIn
}

export default AuthService