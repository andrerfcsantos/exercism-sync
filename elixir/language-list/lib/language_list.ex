defmodule LanguageList do
  def new() do
    []
  end

  def add(list, language) do
    [language | list]
  end

  def remove(list) do
    [_ | tail] = list
    tail
  end

  def first(list) do
    [head | _] = list
    head
  end

  def count([]) do
    0
  end

  def count([_ | tail]) do
    1 + count(tail)
  end

  def functional_list?(list) do
    "Elixir" in list
  end
end
