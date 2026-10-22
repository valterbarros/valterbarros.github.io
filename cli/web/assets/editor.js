window.addEventListener('DOMContentLoaded', () => {
  const easymde = new EasyMDE({
    element: document.getElementById('js-my-text-area'),
  });

  setTimeout(() => {
    console.log('easymde', easymde.value());
  }, 1000)
});
