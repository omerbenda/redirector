async function onFormSubmit(e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const url = formData.get('url');

  const response = await fetch(`?url=${url}`, {
    method: 'POST',
  });

  const hash = await response.text();

  setUrlHash(hash);
}

function setUrlHash(hash) {
  const hashLabel = document.getElementById('hash-label');

  hashLabel.textContent = `${window.location.origin}/${hash}`;
}
