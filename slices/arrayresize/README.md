# Resize eines Arrays über Slices
Diese Demo soll folgende Frage beantworten:
Was passiert, wenn
- ich ein Array habe, auf das zwei Slices zeigen
- ich erweitere das Array über einen der Slices, sodass eine neue Kopie des Arrays entsteht

Jetzt müßte das Array in ein neues Array kopiert werden.

Frage: Zeigen beide Slices jetzt noch auf dasselbe Array, d.h. wurden alle Slice-Zeiger auf das Array aktualisiert?

Ergebnis:
- wenn das Array kopiert werden muss, zeigen die alten Slices immer noch auf das originale Array
- das neue Slice zeigt auf das neue Array
- die Slices modifizieren im Folgenden also unterschiedliche Array
- wenn kein neues Array entsteht, arbeiten die alten und das neue Slice auf dem originalen Array

Fazit: das ist krank! Das führt zu Bugs, die kein Mensch mehr findet.