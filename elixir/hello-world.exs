defmodule HelloWorld do
  def call(msg) do 
    IO.puts msg
  end
end

HelloWorld.call "Hello, World!"