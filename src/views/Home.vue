<script setup>
</script>
<template>
  <html>

  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Not musical instruments</title>
  </head>

  <body>
    <main class="container">
      <h1>Не музыкальные инструменты</h1>
      <p>Добро пожаловать в мир мастерства и творчества! <br>
        Мы рады представить вам наш уникальный магазин инструментов, где каждый может найти идеальное решение для
        своих творческих и профессиональных потребностей. <br>
        У нас вы найдете широкий выбор инструментов высочайшего качества от ведущих мировых производителей.
        Независимо от того, являетесь ли вы профессиональным мастером или начинающим энтузиастом, у нас есть все
        необходимое!

      </p>
      <div id="carousel" class="carousel slide main-carousel" data-bs-ride="carousel">
        <div class="carousel-inner">
          <div class="carousel-item active">
            <img src="assets/pics/pic1.jpeg" class="pic">
          </div>
          <div class="carousel-item">
            <img src="assets/pics/pic2.jpg" class="pic">
          </div>
          <div class="carousel-item">
            <img src="assets/pics/pic3.jpg" class="pic">
          </div>
        </div>
        <button class="carousel-control-prev" type="button" data-bs-target="#carousel" data-bs-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Предыдущий</span>
        </button>
        <button class="carousel-control-next" type="button" data-bs-target="#carousel" data-bs-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Следующий</span>
        </button>
      </div>
      <div>
        <h1>Отзывы о нас</h1>
        <div class="card shadow rounded my-3">
          <div class="card-body">
            <div class="row">
              <div class="col-md-10 d-flex flex-column">
                <div class="product-name">Носов Иван</div>
                <div class="product-description">Отличный магазин! Особенно поразила команда магазина!! Но
                  зачем вам актриса?</div>
              </div>
            </div>
          </div>
        </div>
        <div class="card shadow rounded my-3">
          <div class="card-body">
            <div class="row">
              <div class="col-md-10 d-flex flex-column">
                <div class="product-name">Попов Олег</div>
                <div class="product-description">Ну, выбор так себе.. Молоток, пила и трактер? Может ещё
                  жареных
                  гвоздей добавите?</div>
              </div>
            </div>
          </div>
        </div>
        <div class="card shadow rounded my-3">
          <div class="card-body">
            <div class="row">
              <div class="col-md-10 d-flex flex-column">
                <div class="product-name">Эмма Стоун</div>
                <div class="product-description">Кто там что писал про актрису в команде? Я здесь вообще-то
                  самый первый человек!! </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
    <footer>
    </footer>
  </body>

  </html>
</template>

<script>
export default {
  data() {
    return {
    }
  },
  methods: {
    async testPost() {
      // Отправляем запрос типа POST
      const response = await fetch('/testpost', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: "Иван"
        })
      });

      const jsonResult = await response.json();

      console.log(jsonResult);
    },
    async testGet() {
      // Отправляем запрос типа GET
      const response = await fetch(`/testget?name=Василий`, {
        method: 'GET'
      });

      const textResult = await response.text();

      console.log(textResult);
    }
  }
}
</script>
