import { handleNetworkError } from '.';
import { setUser } from '../actions';

function handleHTTPError(store) {
  return function (response) {
    if (!response.ok) {
      response.json().then((body) => {
        console.log(`response failed with ${response.status}: ${body.Msg}`);
        store.dispatch(setUser({}));
      });
    }
    return response.json();
  };
}

export const loginUser = (store, username, password) => {
  const form = new FormData();
  form.append('username', username);
  form.append('password', password);
  const request = new Request('/login', {
    method: 'post',
    body: form,
    credentials: 'same-origin',
    mode: 'cors',
    redirect: 'follow',
    cache: 'no-cache',
    header: {
      Accept: 'application/json, application/xml, text/plain, text/html, *.*',
      'Content-Type': 'multipart/form-data;',
    },
  });
  fetch(request)
    .then(handleHTTPError(store))
    .then((user) => {
      store.dispatch(setUser(user));
    })
    .catch(handleNetworkError);
};

export const getUser = (store) => {
  const request = new Request('/curuser', {
    method: 'get',
    credentials: 'same-origin',
    mode: 'cors',
    redirect: 'follow',
    cache: 'no-cache',
    header: {
      Accept: 'application/json',
    },
  });
  fetch(request)
    .then(handleHTTPError(store))
    .then((user) => {
      store.dispatch(setUser(user));
    })
    .catch(handleNetworkError);
};

export const logout = (store) => {
  const request = new Request('/logout', {
    method: 'get',
    credentials: 'same-origin',
    mode: 'cors',
    redirect: 'follow',
    cache: 'no-cache',
    header: {
      Accept: 'application/json',
    },
  });
  fetch(request)
    .then(handleHTTPError(store))
    .then((response) => {
      console.log(response);
      store.dispatch(setUser({}));
    })
    .catch(handleNetworkError);
};

