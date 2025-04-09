async function onFormSubmit(e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const url = formData.get('url');

  const response = await fetch(`?url=${url}`, {
    method: 'POST',
  });

  const id = await response.text();

  setUrlId(id);
}

function setUrlId(id) {
  const idLabel = document.getElementById('id-label');

  const href = `${window.location.origin}/${id}`;
  idLabel.href = href;
  idLabel.textContent = href;
}
