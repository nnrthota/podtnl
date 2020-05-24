require "language/go"

class podtnl < Formula
  desc "An interactive terminal based UI application for tracking cryptocurrencies"
  homepage "https://podtnl.sh"
  url "https://github.com/narendranathreddythota/podtnl/archive/1.0.tar.gz"
  sha256 "3b2b039da68c92d597ae4a6a89aab58d9741132efd514bbf5cf1a1a151b16213"
  revision 1
  head "https://github.com/narendranathreddythota/podtnl.git"
  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/narendranathreddythota/podtnl"
    system "go", "get", "-u", "github.com/narendranathreddythota/podtnl"
    cd path do
      system "go", "build", "-o", "#{bin}/podtnl"
    end
  end

  test do
    system "true"
  end
end 