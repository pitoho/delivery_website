<template>
    <div class="product-card">
      <img :src="product.dish_image_path" :alt="product.dish_name" class="product-image" />
      <div class="price-filed">
        <h2>{{ product.dish_name }}</h2>
        <p>{{ product.price }} руб.</p>
      </div>

    </div>
  </template>
  
  <script>
  export default {
    name: 'ProductCard',
    props: {
      product: {
        type: Object,
        required: true
      }
    }
  }
  </script>
  
  <style scoped>
  .product-card {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 16px;
    text-align: center;
    margin: 8px;
    max-width: 200px;
  }
  
  .product-image {
    width: 100%;
    height: auto;
  }
  </style>