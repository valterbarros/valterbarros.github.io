import { EditorState } from "https://esm.sh/@codemirror/state";
import { EditorView, keymap, lineNumbers } from "https://esm.sh/@codemirror/view";
import { markdown } from "https://esm.sh/@codemirror/lang-markdown";
import { defaultKeymap, history, historyKeymap } from "https://esm.sh/@codemirror/commands";
import { syntaxHighlighting, defaultHighlightStyle } from "https://esm.sh/@codemirror/language";// websocket

const ws = new WebSocket("ws://localhost:12769/ws");

const getFileFromUrl = () => {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get('name');
};

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

window.addEventListener('DOMContentLoaded', async () => {
  listenUpdatePreview();
  const sendChange = debounceFn((value) => ws.send(JSON.stringify({ type: 'change', value: value, file: getFileFromUrl() })), 100);

  const updateList = EditorView.updateListener.of((update) => {
    if (update.docChanged) {
      const textoAtual = update.state.doc.toString();
      
      sendChange(textoAtual);
    }
  });

  // TODO: save text

  const temp = document.getElementById("js-data-article");

  const startState = EditorState.create({
    doc: temp.innerHTML,
    extensions: [
      lineNumbers(),
      history(),
      keymap.of([...defaultKeymap, ...historyKeymap]),
      markdown(),
      EditorView.lineWrapping,
      syntaxHighlighting(defaultHighlightStyle),

      EditorView.contentAttributes.of({
        spellcheck: "true",
        lang: "pt-BR"
      }),
      updateList
    ]
  });

  const parent = document.getElementById("editor-container")
  new EditorView({
    state: startState,
    parent
  });

  sendChange(temp.innerHTML);
});
