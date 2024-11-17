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
        card.className = 'col-md-4 mb-3';
        card.innerHTML = `
                    <div class="card">
                        <div class="card-body">
                            <h5>${product.name}</h5>
                            <p>Количество: ${product.quantity}</p>
                            <p>Цена: ${product.price} руб.</p>
                            <button onclick='deleteProduct(${product.id})' data-id="${product.id}"class='btn btn-danger'>Удалить</button> 
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
    body: JSON.stringify({ name, quantity: parseInt(quantity), price: parseFloat(price) })
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
function deleteProduct(button) {
    // Find the closest card element to the button clicked
    const productCard = button.closest('.card'); // Use closest to find the nearest card
    if (productCard) {
        const productId = button.getAttribute('data-id'); // Get the product ID from data attribute

        // Make a DELETE request to the server
        fetch(`/products/${productId}`, { method: 'DELETE' })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Ошибка при удалении товара');
                }
                return response.json(); // Optional, depending on your server response
            })
            .then(() => {
                // Remove the product card from the DOM
                productCard.remove();
            })
            .catch(error => console.error('Ошибка при удалении товара:', error));
    } else {
        console.error('Product card not found');
    }
}
