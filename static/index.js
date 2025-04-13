async function onFormUrlSubmit(e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const url = formData.get('url');

  const response = await fetch('', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ url }),
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

async function onFormUpdateUrlSubmit(e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const id = formData.get('id');
  const url = formData.get('url');

  const response = await fetch('', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id, url }),
  });

  const status = response.status;

  setUpdateStatus(status);
}

function setUpdateStatus(status) {
  const statusLabel = document.getElementById('update-status-label');

  switch (status) {
    case 200:
      statusLabel.textContent = 'Updated successfully!';
      break;
    default:
      statusLabel.textContent = 'Failed to update!';
      break;
  }
}
