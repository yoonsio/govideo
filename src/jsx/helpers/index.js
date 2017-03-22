
export function handleHTTPError() {
  return function (response) {
    if (!response.ok) {
      response.json().then((body) => {
        console.log(`response failed with ${response.status}: ${body.Msg}`);
      });
    }
    return response.json();
  };
}

export function handleNetworkError() {
  return function () {
    // handle network error
  };
}