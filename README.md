# Shortcuts
## GitBash Windows
Vagrant-Verzeichnis ansteuern
````
cd /c/01arbeit/03udemy/golang
````
mygo-Projekt ansteuern
````
cd /c/01arbeit/03udemy/golang/github.com/vkoster/mygo
````

## GitBash VM
mygo Projekt ansteuern
````
cd /vagrant/github.com/vkoster/mygo/
````
# Build-Struktur
Jedes Thema hat ein eigenes Verzeichnis unter /mygo, z.B.:
- mygo/generics
- mygo/packages-demo

Jedes Thema wird von einem zentralen Docker Compose File verwaltet, das sich im Root-Verzeichnis 
des Themas (s.o.) befindet.

Über dieses File kann man:
- Linux App bauen
- Stage 1 alleine bauen
- Multistage Container bauen