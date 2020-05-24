class Podtnl < Formula
  desc "Expose your pod to Online from any kubernetes clusters"
  homepage "https://podtnl.sh"
  url "https://github.com/narendranathreddythota/podtnl/archive/1.0.tar.gz"
  sha256 "65f934bff1199e3e454678a7af776578ccaff82bf3455fcccbc84f1dd62b530a"
  head "https://github.com/narendranathreddythota/podtnl.git"
  depends_on "go" => :build
  def install
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/narendranathreddythota/podtnl"
    path.install Dir["*"]
    cd path do
      system "go", "build", "-o", "#{bin}/podtnl"
    end
  end
  test do
    assert_match version.to_s, shell_output("#{bin}/podtnl", "-v")
  end
end
