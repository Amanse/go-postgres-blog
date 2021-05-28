import Nav from "../components/Navbar";

import "../styles/globals.css";
import "tailwindcss/tailwind.css";

function MyApp({ Component, pageProps }) {
  return (
    <>
      <Nav />

      <main>
        <div className="py-6 mx-auto max-w-7xl sm:px-6 lg:px-8">
          {/* <!-- Replace with your content --> */}
          <div className="flex justify-center sm:px-0">
              <Component {...pageProps} />
          </div>
          {/* <!-- /End replace --> */}
        </div>
      </main>
    </>
  );
}

export default MyApp;
