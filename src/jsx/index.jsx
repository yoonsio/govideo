import React from 'react';
import { render } from 'react-dom';
import configureStore from './store/configureStore';
import App from './App';
import { getUser } from './helpers/userReq';

const store = configureStore();

render(
  <App store={store} />,
  document.getElementById('app'),
);

getUser(store);

/*
console.log(store.getState().user);
// subscribe to redux state change
const unsubscribe = store.subscribe(() => {
  console.log(store.getState().user);
});
unsubscribe();
*/
