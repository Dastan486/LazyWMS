function loadProducts() {
  fetch('/products')
    .then(response => {
      if (!response.ok) {
        throw new Error('Ошибка при загрузке товаров');
      }
      return response.json();
    })
    .then(data => {
      const itemList = document.getElementById('itemList');
      itemList.innerHTML = '';

      data.forEach(product => {
        const card = document.createElement('div');
        card.className = 'col-md-4 mb-3'; // Используем Bootstrap классы для сетки
        card.innerHTML = `
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">${product.name}</h5>
                            <p class="card-text">Количество: ${product.quantity}</p>
                            <p class="card-text">Цена: ${product.price} руб.</p>
                            <button class="btn btn-danger" onclick='deleteProduct(${product.id})'>Удалить</button>
                        </div>
                    </div>`;
        itemList.appendChild(card);
      });
    })
    .catch(error => console.error('Ошибка при загрузке товаров:', error));
}

function addProduct(event) {
  event.preventDefault();

  const name = document.getElementById('itemName').value;
  const quantity = document.getElementById('itemQuantity').value;
  const price = document.getElementById('itemPrice').value;

  fetch('/products', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, quantity: parseInt(quantity), price: parseFloat(price) }) // Преобразование типов
  })
    .then(response => {
      if (!response.ok) {
        return response.json().then(err => { throw new Error(err.error); });
      }
      return response.json();
    })
    .then(() => {
      loadProducts();
      document.getElementById('itemForm').reset();
    })
    .catch(error => console.error('Ошибка при добавлении товара:', error));
}

// Удаление продукта по ID
function deleteProduct(id) {
  fetch(`/products/${id}`, { method: 'DELETE' })
    .then(response => {
      if (!response.ok) {
        throw new Error('Ошибка при удалении товара');
      }
      return response.json();
    })
    .then(() => loadProducts())
    .catch(error => console.error('Ошибка при удалении товара:', error));
}

// Загрузка продуктов при загрузке страницы
document.addEventListener('DOMContentLoaded', loadProducts);

// Функции для работы с поставщиками
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

// Загрузка поставщиков при загрузке страницы на соответствующей странице
if (document.title === "Управление поставщиками") {
  document.addEventListener('DOMContentLoaded', loadSuppliers);
}
