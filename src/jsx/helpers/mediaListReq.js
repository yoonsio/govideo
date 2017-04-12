import { handleHTTPError, handleNetworkError } from '.';

export const getMediaList = (component) => {
  const request = new Request('/listMedia', {
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
    .then(handleHTTPError())
    .then((list) => {
      component.setState({ list });
    })
    .catch(handleNetworkError);
};

