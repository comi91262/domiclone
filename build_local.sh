#! /usr/bin/env bash

composer install
yarn install

cp .env.example .env
php artisan key:generate

yarn run dev

touch database/database.sqlite

php artisan migrate
php artisan db:seed
