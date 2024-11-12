En este proyecto de simulador de estacionamiento en Go, implementé la gestión de vehículos que entran, estacionan y salen en intervalos definidos. Utilicé las características de Go y varias librerías específicas para asegurar un funcionamiento fluido y eficiente. A continuación, explicaré cómo el lenguaje y sus librerías me ayudaron a construir y organizar el simulador.

Go y su Uso en el Proyecto
Go fue ideal para este simulador, ya que requiere manejar concurrencia para que cada vehículo actúe como una entidad independiente. En el simulador, cada vehículo es controlado como una goroutine, lo que permite ejecutar múltiples operaciones en paralelo sin una gran sobrecarga. Esto fue crucial para simular un flujo de entrada y salida en el estacionamiento, donde cada vehículo sigue su propio “reloj” de entrada, estacionamiento y salida.

Ventajas de Go en el Proyecto:

Concurrencia Optimizada: Gracias a las goroutines, pude ejecutar cada vehículo en paralelo, algo esencial en esta simulación.
Sincronización Sencilla: Go proporciona herramientas nativas para manejar el acceso seguro a los carriles de estacionamiento compartidos, lo que me ayudó a evitar problemas de condiciones de carrera.
Manejo Simple del Tiempo: Con la biblioteca time, pude definir tiempos de espera, lo cual facilitó mucho la simulación de entrada y salida de vehículos.
Librerías Utilizadas
sync La librería sync fue clave para la sincronización entre goroutines. Usé sync.Mutex para asegurar que solo un vehículo pudiera acceder y modificar datos compartidos, como la asignación de carriles, en un momento dado. Esto me permitió evitar conflictos al momento de asignar o liberar un carril, manteniendo la integridad de los datos.

time La librería time me permitió manejar los intervalos de entrada, estacionamiento y salida de los vehículos. Usé funciones como time.Sleep() para establecer los tiempos de espera y simular que cada vehículo permanecía estacionado por un tiempo aleatorio antes de salir. Esto dio al simulador un flujo más realista.

math/rand Con math/rand, generé números aleatorios para que cada vehículo tuviera un tiempo de salida diferente. Esto añadió variabilidad al simulador y ayudó a representar condiciones más realistas de un estacionamiento.

github.com/faiface/pixel Para crear la interfaz gráfica, utilicé la librería Pixel, que facilita la creación de gráficos en Go. Pixel me permitió representar visualmente la posición de cada vehículo y los carriles, y fue fundamental para ver cómo se gestionaba la ocupación de espacios en tiempo real.

Organización del Proyecto
Config: Aquí definí constantes globales en const.go, como los límites de tiempo y el tamaño del estacionamiento. Esto me permitió cambiar parámetros del simulador de forma centralizada sin tener que buscar cada referencia en el código.

Models: En types.go definí las estructuras de los vehículos y del estacionamiento. Estas estructuras organizan la información de cada vehículo y su estado en el estacionamiento, lo cual me facilitó manejar los datos a lo largo del simulador.

Services: Separé las operaciones principales en servicios:

Vehicle Services: Incluyen funciones para gestionar la entrada, el estacionamiento y la salida de cada vehículo, controlando la ocupación de carriles y el tiempo de estacionamiento.
Parking Service: Organiza el comportamiento del estacionamiento, manejando el estado de ocupación, asignación de carriles y liberación de espacio cuando un vehículo sale.
Simulator: parking_simulator.go es el motor principal que lanza la simulación. Aquí inicialicé las goroutines de los vehículos y definí los intervalos de entrada y salida. Utilicé time.Sleep y math/rand para simular el tiempo de estacionamiento de cada vehículo y dar al simulador un flujo realista de entrada y salida.

UI: La interfaz gráfica la desarrollé en scene.go y view.go utilizando Pixel. scene.go configura la interfaz y actualiza la representación visual de los carriles y vehículos. view.go define los elementos gráficos, como contenedores y etiquetas, para mostrar en tiempo real el estado del estacionamiento.

Mi Experiencia
Trabajar en este proyecto fue una experiencia enriquecedora. Go resultó ser un lenguaje excelente para tareas concurrentes como esta simulación. Sin embargo, integrarlo con gráficos en tiempo real fue un desafío, ya que Go no tiene soporte gráfico nativo y dependí de una librería externa como Pixel.

Ventajas:

Concurrencia nativa: Go facilita la ejecución de procesos en paralelo, algo que fue esencial para simular vehículos independientes.
Código modular y claro: La estructura de Go permite una organización modular, facilitando la separación de responsabilidades en distintos servicios y modelos.
Desventajas:

Limitaciones en gráficos: Integrar una librería gráfica externa fue complicado, y requirió un tiempo adicional para ajustar el diseño gráfico del proyecto.
Sincronización compleja: Aunque Go proporciona herramientas para la concurrencia, manejar la sincronización entre varias goroutines al mismo tiempo se vuelve desafiante.
Cómo Usarlo
Compilación: Para ejecutar el simulador, uso go run main.go.
Instalación de dependencias: Ejecuté go mod tidy para asegurar que todas las dependencias estén listas y actualizadas para la simulación gráfica.
En general, este proyecto me permitió explorar y entender tanto las fortalezas de Go para la concurrencia como las dificultades de integrar gráficos en tiempo real. Fue una experiencia valiosa, combinando programación de sistemas y visualización gráfica en un simulador que refleja el funcionamiento de un estacionamiento.

Experiencias y Lecciones Aprendidas en el Proyecto de Concurrencia en Go:

Trabajar en este proyecto de concurrencia en Go me ha dado una comprensión profunda sobre cómo diseñar aplicaciones que pueden manejar múltiples tareas de manera simultánea, una habilidad cada vez más importante en el desarrollo de software moderno. La naturaleza concurrente de Go y el manejo de goroutines y canales han transformado mi forma de ver y estructurar aplicaciones que requieren eficiencia y rapidez en la ejecución de tareas.

Una de las lecciones más valiosas ha sido entender y respetar la complejidad de la sincronización entre goroutines. Al principio, manejar el acceso a recursos compartidos y prevenir condiciones de carrera fue todo un reto. Sin embargo, la experiencia me mostró la importancia de herramientas como sync.Mutex para proteger datos y sync.WaitGroup para coordinar la finalización de múltiples goroutines. Estos conceptos son claves para construir software concurrente y me dieron una base sólida para entender mejor la concurrencia en otros lenguajes también.

Además, trabajar con canales en Go me enseñó a construir sistemas en los que las distintas partes de un programa se comunican de manera clara y efectiva. Los canales no solo me ayudaron a reducir errores, sino que también hicieron que el código fuera más fácil de leer y entender, una habilidad crucial en cualquier equipo de desarrollo.

Consejos para Nuevos Programadores en Concurrencia
Empieza por lo básico: Familiarízate bien con las goroutines y cómo lanzar funciones en paralelo. Experimenta con ejemplos simples antes de pasar a aplicaciones complejas.

Usa sync.Mutex y sync.WaitGroup sabiamente: Cuando trabajas con variables compartidas, usa Mutex para evitar problemas de concurrencia. WaitGroup es especialmente útil para asegurarte de que todas las goroutines finalicen antes de que el programa termine.

Aprovecha los canales: Los canales son una de las características más poderosas de Go. Aprender a usarlos para pasar datos entre goroutines puede hacer que tus programas concurrentes sean más seguros y organizados.

Piensa en la estructura del código desde el inicio: Un buen diseño es esencial para evitar problemas al escalar el proyecto. Separar las responsabilidades y agrupar funciones relacionadas en módulos puede hacer que el código sea mucho más manejable.

Prueba y revisa el código concurrente: Los problemas de concurrencia son difíciles de detectar y resolver. Prueba tu código en diferentes escenarios y revisa la lógica para evitar errores sutiles que solo aparecen cuando el programa está bajo carga.

![WhatsApp Image 2024-11-11 at 11 59 08 PM](https://github.com/user-attachments/assets/36ead091-5f34-411e-be3a-ad5be16b61d5)

[Introducción a Go (2).pdf](https://github.com/user-attachments/files/17711833/Introduccion.a.Go.2.pdf)
