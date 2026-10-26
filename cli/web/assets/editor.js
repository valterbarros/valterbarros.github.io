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
    toolbar: [
      'bold', 'italic', 'heading', '|',
      'quote', 'unordered-list', 'ordered-list', '|',
      'link', 'image', '|',
      'fullscreen', '|',
      {
        name: 'save',
        action: function customFunction(editor) {
          const value = editor.value();
          ws.send(JSON.stringify({ type: 'update', value }));
        },
        className: 'fa fa-floppy-o',
        title: 'Save',
      }
    ]
  });

  const send = debounceFn((value) => ws.send(JSON.stringify({ type: 'change', value })), 100);

  send(easymde.value());

  easymde.codemirror.on('change', () => {
    const value = easymde.value();

    send(value);
  });
});
