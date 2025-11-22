package io.github.fatalistix;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicReference;

@SuppressWarnings("unused")
public class WorkStealing {

    public List<Coroutine> steal(AtomicReference<List<Coroutine>> queueRef) {
        while (true) {
            var q = queueRef.get();
            int n = q.size();
            if (n <= 1) {
                return List.of(); // нечего красть
            }

            // Делим пополам: левое остаётся, правое возвращаем
            int mid = n / 2;

            // Делаем копии, чтобы CAS сравнивал действительно разные объекты
            var left = new ArrayList<>(q.subList(0, mid));
            var right = new ArrayList<>(q.subList(mid, n));

            if (queueRef.compareAndSet(q, left)) {
                return right; // украдено
            }
        }
    }

    public Coroutine take(AtomicReference<List<Coroutine>> queueRef) {
        while (true) {
            var q = queueRef.get();
            if (q.isEmpty()) {
                return null;
            }

            var head = q.getFirst();
            var rest = new ArrayList<>(q.subList(1, q.size()));

            if (queueRef.compareAndSet(q, rest)) {
                return head;
            }
        }
    }
}
