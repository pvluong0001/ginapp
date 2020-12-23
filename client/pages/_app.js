import '../styles/globals.css'
import Provider from "react-redux/lib/components/Provider";
import {createStore} from "redux";

function activePostReducer (state = [], action) {
    switch (action.type){
        case "SELECT_POST":
            return action.payload;
        default:
            return state;
    }
}

const store = createStore(activePostReducer, ['Use redux'])

function MyApp({ Component, pageProps }) {
  return (
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
  )
}

export default MyApp
