require "benchmark"

bubble = ->(list) do
	return list if list.size <= 1 # already sorted
  swapped = true
  while swapped do
    swapped = false
    0.upto(list.size-2) do |i|
      if list[i] > list[i+1]
        list[i], list[i+1] = list[i+1], list[i] # swap values
        swapped = true
      end
    end    
  end

  list
end

puts bubble.call([2, 10, 1, 9, 5, 6, 8, 3, 7, 4])

time = Benchmark.bm(500) do |x|
	x.report(:bubble) { bubble.call([2, 10, 1, 9, 5, 6, 8, 3, 7, 4]) }
end

puts time