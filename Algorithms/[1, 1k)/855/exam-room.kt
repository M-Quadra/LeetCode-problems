class ExamRoom(private val n: Int) {

    private class Line(
        val l: Int, val r: Int, val hasL: Boolean = true, val hasR: Boolean = true
    )

    private val pointSet = TreeSet<Int>()
    private val lineQueue = PriorityQueue<Line> { a, b ->
        val lA = (a.r - a.l) / if (a.hasL && a.hasR) 2 else 1
        val lB = (b.r - b.l) / if (b.hasL && b.hasR) 2 else 1
        if (lA != lB) lB - lA else a.l - b.l
    }.apply {
        add(Line(0, n - 1, hasL = false, hasR = false))
    }

    fun seat(): Int {
        val line = run {
            while (true) {
                val top = lineQueue.poll()!!
                if (pointSet.contains(top.l) != top.hasL || pointSet.contains(top.r) != top.hasR) continue
                if ((pointSet.higher(top.l) ?: top.r) != top.r) continue
                if ((pointSet.lower(top.r) ?: top.l) != top.l) continue
                return@run top
            }
        } as Line

        val p = if (!line.hasL) line.l else if (!line.hasR) line.r else (line.l + line.r) / 2
        pointSet.add(p)
        if (line.l + (if (line.hasL) 1 else 0) < p) lineQueue.add(Line(line.l, p))
        if (p + (if (line.hasR) 1 else 0) < line.r) lineQueue.add(Line(p, line.r, hasR = line.hasR))
        return p
    }

    fun leave(p: Int) {
        val l = pointSet.lower(p) ?: 0
        val r = pointSet.higher(p) ?: (n - 1)
        pointSet.remove(p)
        lineQueue.add(Line(l, r, pointSet.contains(l), pointSet.contains(r)))
    }
}