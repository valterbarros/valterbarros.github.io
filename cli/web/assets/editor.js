// websocket
const ws = new WebSocket("ws://localhost:12769/ws");

const debounceFn = (fn, delay) => {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      fn.apply(this, args);
    }, delay);
  };
};

function listenUpdatePreview() {
  ws.onmessage = function(event) {
    const preview = document.getElementById('js-preview');
    preview.innerHTML = event.data;
  };
}

window.addEventListener('DOMContentLoaded', () => {
  listenUpdatePreview();

  const easymde = new EasyMDE({
    element: document.getElementById('js-my-text-area'),
  });

  const send = debounceFn((value) => ws.send(value), 100);

  easymde.codemirror.on('change', () => {
    const value = easymde.value();

    send(value);
  });
});
