const body = document.body;
const darkModeToggleIcon = document.querySelector('#darkModeToggleIcon');

function toggleTheme() {
  body.classList.toggle('darkMode');
  darkModeToggleIcon.innerHTML = body.classList.contains('darkMode') ? 'dark_mode' : 'light_mode';
  darkModeToggleIcon.style.color = body.classList.contains('darkMode') ? '#0d1117' : 'white';
}

document.querySelector('#uploadForm').addEventListener('submit', async (e) => {
  e.preventDefault();

  const file = document.querySelector('#uploadFile').files[0];
  const formData = new FormData();

  formData.append('file', file);
  await fetch('/files', {
    method: 'PUT',
    body: formData,
  });

  window.location.reload();
});
