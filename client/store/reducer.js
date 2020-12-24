import authReducer from "../domains/auth/reducers/auth";
import {combineReducers} from "redux";

const rootReducer = combineReducers({
    auth: authReducer
})

export default rootReducer