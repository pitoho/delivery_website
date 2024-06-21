import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useBucketStore = defineStore('bucket', () => {

    fetch('http://localhost:3000/dishes')
    .then(res => res.json())
    .then(json => {
        json.map(elem => {
            let res = {
                id_dish: elem.id,
                dish_name: elem.dish_name,
                dish_image_path: elem.dish_image_path,
                price: elem.price,
                tags_id: elem.tags_id
            }
            bucket.value.push(res)
        })
    })  

  let bucket = ref([])
   let order = ref([])
   let orderCost = ref(0)
   const randomFood = localStorage.getItem('randomFood');
   if (randomFood) {
     order.value = JSON.parse(randomFood);
     console.log(bucket.value)
   }
  let buckLength = ref(bucket.value.length)
  let BuckExist = ref(bucket.value.length); 
  let totPrice = ref(0);
   bucket.value.forEach(element => {
       totPrice.value+=element.lastPrice
   });
   reCostOrder()


   function addToBucket(id, title, lastPrice, oldPrice, image, count){
    let findIndex = bucket.value.findIndex(elem => elem.id === id)
    if (findIndex !== -1){
        
        // localStorage.bucket.value = JSON.stringify(bucket.value)
        bucket.value.splice(findIndex, 1)
        buckLength.value--
        localStorage.setItem('randomFood', JSON.stringify(bucket.value))
        console.log(bucket.value)
        // console.log(localStorage.getItem(JSON.stringify(('bucket.value'))))
        // console.log(bucket.value)
    }else{
        bucket.value.push({id: id, title: title, lastPrice: lastPrice, oldPrice: oldPrice, count: count,
            image: image })
        buckLength.value++
            localStorage.randomFood = JSON.stringify(bucket.value)
            console.log(bucket.value)
        //    console.log(localStorage.getItem(JSON.stringify('bucket.value')))
    }
    totPrice.value=0
    bucket.value.forEach(element => {
    totPrice.value+=element.lastPrice
    // BuckExist.value = BuckExist.value - 1
    });
    // if (findIndex !== -1){
    //     bucket.value = bucket.value.filter(elem => elem.id !== id)
    // }else{
    //     bucket.value.push({id: id, title: title, price: price,
    //         image: image, rating: rating, discount: discount, is_available: is_available})
    //        console.log(bucket.value)
    // }

    // bucket.value.push({id: id, title: title, price: price,
    //     image: image, rating: rating, discount: discount, is_available: is_available})
    //    console.log(bucket.value)
}


   function deleteItem(id){
    order.value = order.value.filter(e => e.id_dish != id)
    console.log(id)
    console.log(order.value)
    localStorage.randomFood = JSON.stringify(order.value)
    reCostOrder()
   }

   function addRandomItem(tag){
    let neededProducts = bucket.value.filter(elem => elem.tags_id == tag)
    let item = neededProducts[Math.floor(Math.random()*neededProducts.length)];
    console.log('Added product with tag_id ' + tag + ' and id of product ' + item.id_dish)
    order.value.push(item)
    localStorage.randomFood = JSON.stringify(order.value)
    reCostOrder()
    }

    function reCostOrder(){
        orderCost.value = 0
        for (let item of order.value){
            console.log(item)
            orderCost.value += item.price
        }
    }

  return { bucket, buckLength, BuckExist, totPrice, order, deleteItem, addToBucket, addRandomItem, orderCost, reCostOrder }
})
