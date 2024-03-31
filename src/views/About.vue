<script setup>
</script>

<template>
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>О нас</title>
</head>

<body>
    <main class="container">
        <h2>Наша Миссия</h2>
        <p>Мы стремимся создать место, где каждый, кто разделяет любовь к ремеслу и творчеству, может найти надежные и
            профессиональные инструменты для своих задач.
            Наша миссия — стать вашим надежным партнером в творчестве и строительстве.</p>
        <h2>Наш Подход</h2>
        <p>Мы ценим индивидуальность каждого клиента и понимаем, что ваши потребности уникальны.
            Поэтому мы сосредотачиваемся на предоставлении широкого ассортимента инструментов высшего качества, чтобы
            удовлетворить самые разнообразные потребности.</p>
        <h2>Команда Экспертов</h2>
        <p>Наша команда — это объединение опытных специалистов, готовых поделиться своими знаниями и помочь вам выбрать
            правильный инструмент.
            Мы приветствуем вас в нашем мире, где каждый член команды стремится к общей цели — вашему успеху.</p>
        <div class="row">
            <router-link class="nav-link" to="worker1">Жопа</router-link>
            <div class="col-md-4 worker" onclick="location.href='workers/worker1.html'">
                
                <div class="card shadow rounded my-3">
                    
                    <div class="card-body">
                        <div><img src="assets/pics/worker1.jpeg" class="avatar"></div>
                        <div class="avatar-post bi bi-emoji-sunglasses">Финансовый менеджер</div>
                        <div class="avatar-text">Сделает все деньги мира</div>
                        <div class="avatar-text bi bi-envelope"><a href="mailto:tinkoff@example.ru">Напишите мне!</a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-md-4 worker" onclick="location.href='workers/worker2.html'">
                <div class="card shadow rounded my-3">
                    <div class="card-body">
                        <div><img src="assets/pics/worker2.jpeg" class="avatar"></div>
                        <div class="avatar-post bi bi-star">Актриса</div>
                        <div class="avatar-text">Кто её нанял в магазин инструментов?</div>
                        <div class="avatar-text bi bi-envelope"><a href="mailto:stone@example.ru">Напишите мне!</a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-md-4 worker" onclick="location.href='workers/worker3.html'">
                <div class="card shadow rounded my-3">
                    <div class="card-body">
                        <div><img src="assets/pics/worker3.jpeg" class="avatar"></div>
                        <div class="avatar-post bi bi-thermometer-sun">Маркетинговый директор</div>
                        <div class="avatar-text">Не только фотографируется в барханах</div>
                        <div class="avatar-text bi bi-envelope"><a href="mailto:durov@example.ru">Напишите мне!</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </main>
    <footer>
    </footer>
</body>
</template>

<script>
</script>
