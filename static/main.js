// main.js

// Импортируем функции из других файлов (если используете ES6 модули)
// import { loadProducts, addProduct, deleteProduct } from './products.js';
// import { loadSuppliers, addSupplier } from './suppliers.js';

// Если вы не используете ES6 модули, просто убедитесь, что файлы подключены в HTML

// Загрузка поставщиков при загрузке страницы на соответствующей странице
if (document.title === "Управление поставщиками") {
  document.addEventListener('DOMContentLoaded', loadSuppliers);
}

