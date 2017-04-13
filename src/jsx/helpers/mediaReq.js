import { handleHTTPError, handleNetworkError } from '.';

export const getMedia = (component, encodedPath) => {
  const request = new Request(`/media/${encodedPath}/info`, {
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
    .then((media) => {
      component.setState({ 
        media: media,
        path: `/media/${encodedPath}/data`
      });
      if (media.subtitle != "") {
        component.setState({
          subtitle_path: `/media/${encodedPath}/subtitle`
        });
      }
    })
    .catch(handleNetworkError);
};
