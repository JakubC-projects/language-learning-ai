import  { Auth0Client,User,createAuth0Client } from "@auth0/auth0-spa-js";
import AuthConfig from "../config/auth.json"


export let auth0Client: Auth0Client;

async function configureClient () {
    auth0Client = await createAuth0Client({
      domain: AuthConfig.domain,
      clientId: AuthConfig.clientId,
      authorizationParams: {
        redirect_uri: "http://localhost:8080",
        audience: AuthConfig.audience,
        scope: "profile",
      }
    });
};

export async function getUser(): Promise<User> {
  const user = await auth0Client.getUser();
  if (!user) {
    throw Error("cannot get user")
  }
  return user;
}

export async function ensureIsLoggedIn(): Promise<boolean> { 
    await  configureClient()
    const isAuthenticated = await auth0Client.isAuthenticated();

    if (isAuthenticated) {
      return true;
    }
  
    const query = window.location.search;
    if (query.includes("code=") && query.includes("state=")) {
      await auth0Client.handleRedirectCallback();
      window.history.replaceState({}, document.title, "/");
      return true;
    }

    await auth0Client.loginWithRedirect()

    return false

}

export async function logout(): Promise<void> {
  if (!auth0Client) return;
  return auth0Client.logout({
    logoutParams: {
      returnTo: "http://localhost:8080"
    }
  });
}