<?php

require __DIR__.'/../vendor/autoload.php';

// Begin Generation Logic

$languageName = 'a language';
$description = 'a description of the language';

// Begin Render Logic

$loader = new Twig_Loader_Filesystem( __DIR__.'/../templates' );
$twig = new Twig_Environment( $loader );

$template = $twig->load( 'index.html.twig' );

echo $template->render( [ 'name' => $languageName, 'description' => $description ] );