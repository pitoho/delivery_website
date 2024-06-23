<script setup>
import { useBucketStore } from '@/stores/BucketStore';
import ProductItem from '@/components/ProductItem.vue';
import { ref, watch } from 'vue';
import axios from 'axios';

const bucketStore = useBucketStore()

const error = ref('')
let street = ref('')
let house = ref()
let corpus_building = ref()
let flat = ref()
let orderCost = ref(bucketStore.orderCost)
watch(()=>bucketStore.orderCost,
            (newPrice) => {
				orderCost.value = newPrice
                console.log(orderCost.value)
            })


const order = async () => {
  	try {
    	const response = await axios.post('/order', {
			orderedFood: localStorage.getItem('randomFood'),
      		street: street.value,
      		house: house.value,
	  		corpus_building: corpus_building.value,
	  		flat: flat.value,
	  		totalPrice: orderCost.value,

    	},
		{
      		headers: { 'Content-Type': 'application/json' }
    	});
		console.log(response.data)
		if (response.data.success) {
			console.log('success')
      		const data = await response.data;
      		error.value = data.message || 'Успешный вход'
        	window.location.href = '/success'; 
			// bucketStore.clearBucket()
    	}

  	} catch (err) {
    	error.value = 'Произошла ошибка при отправке запроса';
    	console.error(err);
  	}
}
</script>
 
<template >
<section id="order" class="story">
		<div class="container">
			<h1 class="story__title">Основываясь на ваших предпочтениях:</h1>
			<p class="offer__intro">(дважды кликните по товару, если он Вам не понравился)</p>
		</div>
		<div class="story-content container container_full-height" style="flex-direction: column;" id ="story-content">
			<div class="story-content container container_full-height forScroll" style="padding: 0px;" id ="story-content">
				<ProductItem v-for="product in bucketStore.order" :key="product.id_dish" :product="product"/>
			</div>
				<div class="adressInputText" style="width: 100%; padding: 40px;">Сумма заказа:   <b style="font-size: 24px;">{{  orderCost }} Р</b></div>
		</div>
		<form @submit.prevent="order">
			<div class="adressField">
				<label for="email" class="adressInputText">Выберите адрес для доставки:</label>
			</div>
			<div class="adressField">
				<label for="email" class="adressInputText">Улица:</label>
				<input style="width: 300px;" class="adressInput" type="text" id="email" v-model="street" required>
				<label for="email" class="adressInputText">Дом:</label>
				<input style="width: 100px;" class="adressInput" type='number' id="email" v-model="house" required>
				<label for="email" class="adressInputText">Корпус:</label>
				<input style="width: 100px;" class="adressInput" type='number' id="email" v-model="corpus_building">
				<label for="email" class="adressInputText">Квартира:</label>
				<input style="width: 100px;" class="adressInput" type='number' id="email" v-model="flat" required>
			</div>
			<button class="offer__btn btn">Заказать</button>
		</form>
</section>
</template>

<style scoped>
.adressInputText{
	font-size: 16px;
}
.adressInput{
	padding: 12px 0;
	font-size: 16px;
	font-weight: 800;
	display: block;
	/* width: 500px; */
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
	text-align: center;
}
.adressField{
	display: flex;
	justify-content: center;
	align-items: center;
	gap: 20px;
	margin: 20px;
}
button{
	border: none;
}
.btn {
	padding: 18px 0;
	background: #FF3773;
	color: #fff;
	font-size: 16px;
	font-weight: 800;
	display: block;
	width: 244px;
	margin: 0 auto;
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
	text-align: center;
}
.btn:hover {
	background: #FF185E;
	transform: scale(1.5,1.5);
}

.btn-ghost {
	padding: 18px 0;
	color: #FF3773;
	border: 1px solid #FF3773;
	font-size: 16px;
	font-weight: 800;
	display: block;
	width: 329px;
	text-align: center;
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
}

.btn-ghost:hover {
	color: #fff;
	background: #FF3773;
	border: 1px solid #FF3773;
}

    .story-content {
	padding: 20px;
	background: #F1F1F5 no-repeat right / 55%;
	height: fit-content;
	display: flex;
}

.forScroll{
	overflow-x: scroll;
}

.story-content__item {
	font-size: 16px;
	width: 401px;
	position: relative;
	padding-left: 22px;
	line-height: 1.01em;
	margin-bottom: 74px;
	line-height:20px
}

.story-content__item:before {
	content:'';
	display: block;
	width: 7px;
	height: 40px;
	background: #9E9EB7;
	position: absolute;
	left: 0;
}
.story__title {
	font-size: 60px;
	letter-spacing: 0.22em;
	line-height: 1.35em;
	background-color: rgba(196, 210, 247, 0);
}
.story{
    padding-top: 5vh;
		/* height: 130vh; */
	}
    .story-content__list{
	margin: 0 auto;
}

</style>