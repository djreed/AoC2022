The crux of this problem is creating a comparator that returns the relative greater-or-lesser result of two packets.
Using this, we can find which packet pairs are ordered, and count those indices.

For part 2, rather than dealing with sorting the whole list, given we just need to find the locations of the
separators, we can move the separators in the imaginary list of results based on the compared results as we go,
effectively binning input data entries as either behind sep1 or between sep1 and sep2. Entries after sep2 won't
impact the index of sep2 itself, so they can be discarded.