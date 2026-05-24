const form = document.getElementById('hello-form');
const nameInput = document.getElementById('name');
const result = document.getElementById('result');
const responseBox = document.getElementById('response');

const setStatus = (message, isError = false) => {
  result.textContent = message;
  result.classList.toggle('error', isError);
};

form.addEventListener('submit', async (event) => {
  event.preventDefault();
  const name = nameInput.value.trim();

  if (!name) {
    setStatus('Informe um nome para continuar.', true);
    responseBox.textContent = '';
    return;
  }

  try {
    setStatus('Carregando...');
    const res = await fetch(`/api/hello?name=${encodeURIComponent(name)}`);
    const data = await res.json();
    responseBox.textContent = JSON.stringify(data, null, 2);

    if (!res.ok) {
      setStatus(data.error || 'Erro inesperado.', true);
      return;
    }

    setStatus(data.message || 'Resposta recebida.');
  } catch (err) {
    setStatus('Falha ao conectar com o backend.', true);
    responseBox.textContent = '';
  }
});
