// suppliers.js

function loadSuppliers() {
  fetch('/suppliers')
    .then(response => {
      if (!response.ok) {
        throw new Error('Ошибка при загрузке поставщиков');
      }
      return response.json();
    })
    .then(data => {
      const supplierList = document.getElementById('supplierList');
      supplierList.innerHTML = '';

      data.forEach(supplier => {
        const card = document.createElement('div');
        card.className = 'col-md-4 mb-3'; // Используем Bootstrap классы для сетки
        card.innerHTML = `
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">${supplier.name}</h5>
                            <p class="card-text">Контактная информация: ${supplier.contact_info}</p>
                        </div>
                    </div>`;
        supplierList.appendChild(card);
      });
    })
    .catch(error => console.error('Ошибка при загрузке поставщиков:', error));
}

function addSupplier(event) {
  event.preventDefault();

  const name = document.getElementById('supplierName').value;
  const contactInfo = document.getElementById('supplierContactInfo').value;

  fetch('/suppliers', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, contact_info: contactInfo })
  })
    .then(response => {
      if (!response.ok) {
        return response.json().then(err => { throw new Error(err.error); });
      }
      return response.json();
    })
    .then(() => {
      loadSuppliers();
      document.getElementById('supplierForm').reset();
    })
    .catch(error => console.error('Ошибка при добавлении поставщика:', error));
}

// Загрузка поставщиков при загрузке страницы
document.addEventListener('DOMContentLoaded', loadSuppliers);

