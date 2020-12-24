import '../styles/globals.css'

import rootStore from "../store";
import {Provider} from "react-redux";
function MyApp({ Component, pageProps }) {
  return (
      <Provider store={rootStore}>
        <Component {...pageProps} />
      </Provider>
  )
}

export default MyApp
