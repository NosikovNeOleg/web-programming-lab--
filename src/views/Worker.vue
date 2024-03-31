<script setup>
</script>

<template>
<body>
    <main class="container">
        <div class="row">
            <div class="col-md-4">
                <div><img src="assets/pics/worker1.jpeg" class="avatar"></div>
                <div class="avatar-post">Олег Тиньков</div>
                <div class="avatar-text">Финансовый менеджер</div>
                <div class="avatar-text bi bi-envelope"><a href="mailto:stone@example.ru">Напишите мне!</a></div>
            </div>
            <div class="col-md-4 text-about">
                <h2>Моя Миссия</h2>
                <p>Мы понимаем, что в мире финансов важным является не только результат, но и процесс достижения целей.
                    В нашем магазине представлены инструменты, обеспечивающие безопасность и удобство в управлении
                    вашими финансами. От инновационных технологий до простоты использования — мы создали для вас
                    идеальное финансовое пространство.</p>
            </div>
            <div class="col-md-4 text-about">
                <h2>Мой Подход</h2>
                <p>Я предлагаю вам не только инструменты, но и эксклюзивные бонусы и привилегии. С нами каждая сделка
                    приносит вам не только прибыль, но и дополнительные возможности для роста. Ваш финансовый опыт
                    становится особенным, когда вы выбираете наши инструменты.</p>
            </div>
        </div>

    </main>
    <footer>
    </footer>
</body>
</template>

<script>
export default {
  data() {
    return {
      currentRating: 0,
    }
  },
  computed: { 
  },
  methods: {
  },
  mounted() {
    console.log('Компонент примонтирован!');
  }
}
</script>