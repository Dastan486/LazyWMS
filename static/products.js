// products.js

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

